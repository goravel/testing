package testing

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"goravel/bootstrap"

	"github.com/goravel/framework/support/file"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/filesystem"
	"github.com/stretchr/testify/assert"
)

func TestFilesystem(t *testing.T) {
	bootstrap.Boot()

	type Disk struct {
		disk string
		url  string
	}
	disks := []Disk{
		{
			disk: "local",
			url:  "http://localhost/storage/",
		},
		{
			disk: "oss",
			url:  "https://goravel.oss-cn-beijing.aliyuncs.com/",
		},
		{
			disk: "cos",
			url:  "https://goravel-1257814968.cos.ap-beijing.myqcloud.com/",
		},
		{
			disk: "s3",
			url:  "https://goravel.s3.us-east-2.amazonaws.com/",
		},
		{
			disk: "custom",
			url:  "http://localhost/storage/",
		},
	}

	tests := []struct {
		name  string
		setup func(name string, disk Disk)
	}{
		{
			name: "Put",
			setup: func(name string, disk Disk) {
				assert.Nil(t, facades.Storage.Disk(disk.disk).Put("Put/1.txt", "Goravel"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists("Put/1.txt"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Missing("Put/2.txt"), name)
			},
		},
		{
			name: "Get",
			setup: func(name string, disk Disk) {
				assert.Nil(t, facades.Storage.Disk(disk.disk).Put("Get/1.txt", "Goravel"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists("Get/1.txt"), name)
				data, err := facades.Storage.Disk(disk.disk).Get("Get/1.txt")
				assert.Nil(t, err, name)
				assert.Equal(t, "Goravel", data, name)
				length, err := facades.Storage.Disk(disk.disk).Size("Get/1.txt")
				assert.Nil(t, err, name)
				assert.Equal(t, int64(7), length, name)
			},
		},
		{
			name: "PutFile_Text",
			setup: func(name string, disk Disk) {
				fileInfo, err := filesystem.NewFile("./resources/filesystems/test.txt")
				assert.Nil(t, err, name)
				path, err := facades.Storage.Disk(disk.disk).PutFile("PutFile", fileInfo)
				assert.Nil(t, err, name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists(path), name)
				data, err := facades.Storage.Disk(disk.disk).Get(path)
				assert.Nil(t, err, name)
				assert.Equal(t, "Goravel", data, name)
			},
		},
		{
			name: "PutFile_Image",
			setup: func(name string, disk Disk) {
				fileInfo, err := filesystem.NewFile("./resources/logo.png")
				assert.Nil(t, err, name)
				path, err := facades.Storage.Disk(disk.disk).PutFile("PutFile", fileInfo)
				assert.Nil(t, err, name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists(path), name)
			},
		},
		{
			name: "PutFileAs_Text",
			setup: func(name string, disk Disk) {
				fileInfo, err := filesystem.NewFile("./resources/filesystems/test.txt")
				assert.Nil(t, err, name)
				path, err := facades.Storage.Disk(disk.disk).PutFileAs("PutFileAs", fileInfo, "text")
				assert.Nil(t, err, name)
				assert.Equal(t, "PutFileAs/text.txt", path, name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists(path), name)
				data, err := facades.Storage.Disk(disk.disk).Get(path)
				assert.Nil(t, err, name)
				assert.Equal(t, "Goravel", data, name)

				path, err = facades.Storage.Disk(disk.disk).PutFileAs("PutFileAs", fileInfo, "text1.txt")
				assert.Nil(t, err, name)
				assert.Equal(t, "PutFileAs/text1.txt", path, name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists(path), name)
				data, err = facades.Storage.Disk(disk.disk).Get(path)
				assert.Nil(t, err, name)
				assert.Equal(t, "Goravel", data, name)
			},
		},
		{
			name: "PutFileAs_Image",
			setup: func(name string, disk Disk) {
				fileInfo, err := filesystem.NewFile("./resources/logo.png")
				assert.Nil(t, err, name)
				path, err := facades.Storage.Disk(disk.disk).PutFileAs("PutFileAs", fileInfo, "image")
				assert.Nil(t, err, name)
				assert.Equal(t, "PutFileAs/image.png", path, name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists(path), name)

				path, err = facades.Storage.Disk(disk.disk).PutFileAs("PutFileAs", fileInfo, "image1.png")
				assert.Nil(t, err, name)
				assert.Equal(t, "PutFileAs/image1.png", path, name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists(path), name)
			},
		},
		{
			name: "Url",
			setup: func(name string, disk Disk) {
				assert.Nil(t, facades.Storage.Disk(disk.disk).Put("Url/1.txt", "Goravel"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists("Url/1.txt"), name)
				assert.Equal(t, disk.url+"Url/1.txt", facades.Storage.Disk(disk.disk).Url("Url/1.txt"), name)
				if disk.disk != "local" && disk.disk != "custom" {
					resp, err := http.Get(disk.url + "Url/1.txt")
					assert.Nil(t, err, name)
					content, err := ioutil.ReadAll(resp.Body)
					resp.Body.Close()
					assert.Nil(t, err, name)
					assert.Equal(t, "Goravel", string(content), name)
				}
			},
		},
		{
			name: "TemporaryUrl",
			setup: func(name string, disk Disk) {
				assert.Nil(t, facades.Storage.Disk(disk.disk).Put("TemporaryUrl/1.txt", "Goravel"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists("TemporaryUrl/1.txt"), name)
				url, err := facades.Storage.Disk(disk.disk).TemporaryUrl("TemporaryUrl/1.txt", time.Now().Add(5*time.Second))
				assert.Nil(t, err, name)
				assert.NotEmpty(t, url, name)
				if disk.disk != "local" && disk.disk != "custom" {
					resp, err := http.Get(url)
					assert.Nil(t, err, name)
					content, err := ioutil.ReadAll(resp.Body)
					resp.Body.Close()
					assert.Nil(t, err, name)
					assert.Equal(t, "Goravel", string(content), name)
				}
			},
		},
		{
			name: "Copy",
			setup: func(name string, disk Disk) {
				assert.Nil(t, facades.Storage.Disk(disk.disk).Put("Copy/1.txt", "Goravel"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists("Copy/1.txt"), name)
				assert.Nil(t, facades.Storage.Disk(disk.disk).Copy("Copy/1.txt", "Copy1/1.txt"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists("Copy/1.txt"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists("Copy1/1.txt"), name)
			},
		},
		{
			name: "Move",
			setup: func(name string, disk Disk) {
				assert.Nil(t, facades.Storage.Disk(disk.disk).Put("Move/1.txt", "Goravel"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists("Move/1.txt"), name)
				assert.Nil(t, facades.Storage.Disk(disk.disk).Move("Move/1.txt", "Move1/1.txt"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Missing("Move/1.txt"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists("Move1/1.txt"), name)
			},
		},
		{
			name: "Delete",
			setup: func(name string, disk Disk) {
				assert.Nil(t, facades.Storage.Disk(disk.disk).Put("Delete/1.txt", "Goravel"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists("Delete/1.txt"), name)
				assert.Nil(t, facades.Storage.Disk(disk.disk).Delete("Delete/1.txt"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Missing("Delete/1.txt"), name)
			},
		},
		{
			name: "MakeDirectory",
			setup: func(name string, disk Disk) {
				assert.Nil(t, facades.Storage.Disk(disk.disk).MakeDirectory("MakeDirectory1/"), name)
				assert.Nil(t, facades.Storage.Disk(disk.disk).MakeDirectory("MakeDirectory2"), name)
				assert.Nil(t, facades.Storage.Disk(disk.disk).MakeDirectory("MakeDirectory3/MakeDirectory4"), name)
			},
		},
		{
			name: "DeleteDirectory",
			setup: func(name string, disk Disk) {
				assert.Nil(t, facades.Storage.Disk(disk.disk).Put("DeleteDirectory/1.txt", "Goravel"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Exists("DeleteDirectory/1.txt"), name)
				assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("DeleteDirectory"), name)
				assert.True(t, facades.Storage.Disk(disk.disk).Missing("DeleteDirectory/1.txt"), name)
			},
		},
	}

	for _, disk := range disks {
		for _, test := range tests {
			test.setup(disk.disk+" "+test.name, disk)
		}

		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("Put"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("Get"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("PutFile"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("PutFileAs"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("Url"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("TemporaryUrl"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("Copy"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("Copy1"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("Move"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("Move1"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("Delete"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("MakeDirectory1"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("MakeDirectory2"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("MakeDirectory3"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("MakeDirectory4"), disk.disk)
		assert.Nil(t, facades.Storage.Disk(disk.disk).DeleteDirectory("DeleteDirectory"), disk.disk)

		if disk.disk == "local" || disk.disk == "custom" {
			assert.True(t, file.Remove("./storage"))
		}
	}
}
