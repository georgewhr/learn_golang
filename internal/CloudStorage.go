package internal

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
In memory

Level 1:
adding, get, copy files

["COPY_FILE", "/not-existing.file", "/dirl/file.txt"],
"/dirl/dir2/file.txt", "/dirl/file.txt"],
"/dir1/file.txt", "15"],
["COPY_FILE", "/dirl/file.txt", "/dirl/dir2/file.txt"],
returns "10"
Explanations
queries = [
["ADD_FILE", "/dir1/dir2/file.txt", "10"1,
["COPY_FILE", ["ADD_FILE",
["GET_FILE_SIZE", "/dirl/file. txt"],
["GET_FILE_SIZE", "/not-existing.]




Level 2:
matchhing prefix and surfix
Implement support for retrieving file names by searching directories via prefixes and suffixes.
• FIND_FILE < prefix› ‹ suffix> - should search for files with names starting with prefix and ending with suffix. Returns a string representing all matching files in this format: "‹namel> (<sizel>), <name2> (<size2>), ..." . The output should be sorted in descending order of file sizes or, in the case of ties, lexicographically. If no files match the required properties, should return an empty string.
queries = [
["ADD_FILE", "/root/dir/another_dix/file.mp3", "10"],
["ADD_FILE", "/root/file.mp3", "5"],|
["ADD_FILE", "/root/musie/file.mp3", "7"],
['COPY_FILE", "/root/music/file.mp3", "/root/dix/file.mp3"],
("FIND_FILE", "/root", ".mp3"],
("FIND_FILE", "/root", "file.txt"],
["FIND_FILE", "/dir", "file.mp3"]
returns "true" returns "true" returns "true" returns "true"
returns "/root/dir/another_dix/file.mp3(10), /root/diz/file.mp3(7), /root/music/file.mp3 (7) , /root/file.mp3 (5) "| returns ""; there is no file with the prefix "/root" and suffix "file.txt" returns ""; there is no file with the prefix "/dir" and suffix "file.mp3"

Level 3:
add users and capacity limits
Implement support for different users sending queries to the system. All users share a common filesystem in the cloud storage, but each user is assigned an individual storage capacity limit.
• ADD_USER <userId> <capacity> — Should add a new user to the system, with capacity as their storage limit in bytes. The total size of all files owned by userid cannot exceed capacity . The operation fails if a user with userid already exists. Returns "true" if a user with userid is successfully created, or "false" otherwise.
• ADD_FILE_BY <userId> ‹name> ‹size> — Should behave in the same way as the ADD_FILE from Level 1, but the added file should be owned by the user with userid. A new file cannot be added to the storage if doing so will exceed the user's capacity limit. Returns a string representing the remaining storage capacity for userid if the file is successfully added or an empty string otherwise.
Note that all queries calling the ADD_FILE operation implemented during Level 1 are run by the user with userid = "admin", who has unlimited storage capacity. Also, assume that the COPY_FILE operation
preserves the ownership of the original file.
• UPDATE_CAPACITY <userId> ‹capacity> - Should change the maximum storage capacity for the user with userid. If the total size of all user's files exceeds the new capacity, the largest files (sorted lexicographically in case of a tie) should be removed from the storage until the total size of all remaining files will no longer exceed the new capacity. Returns a string representing the number of removed files, or an empty string if a user with userid does not

["ADD_USER", "userl", "125"], ['ADD_USER", "userl", "100"], ['ADD_USER", "user2", "100"],
["ADD_FILE_BY", "userl", "/dir/file.big", "50"],
["ADD_FILE_BY", "userl",
["ADD_FILE_BY",
"user2", "/file.med", "40"],
["COPY_FILE", "/file.med", "/dir/another/file.med"],
["COPY_FILE", "/file.med", "/dir/another/another/file.med"],
["ADD_FILE_BY", "userl", "/dir/file small", "10"1,
["ADD_FILE", "/dir/admin_file",
["ADD_FILE_BY", "userl",
"/dir/file.small", "5"],
["ADD_FILE_BY", "userl",
"/my_folder/file. huge",
"100"],
["ADD_FILE_BY", "user3", "/my_folder/file.huge", "100"l,
["UPDATE_CAPACITY",
"userl", "300"l,
["UPDATE_CAPACITY", "userl", "50"],
returns "true"; creates user "userl" with 125 bytes capacity returns "false"; "userl" already exists
returns "true"; creates user "user2" with 100 bytes capacity
['"UPDATE_CAPACITY",
"user2",

returns "true"; creates user "userl" with 125 bytes capacity returns "false"; "userl" already exists
returns "true"; creates user "user2" with 100 bytes capacity
['"UPDATE_CAPACITY",
"user2", "1000"]
returns ""; file named "/file.med" already exists and owned by "userl"
returns "true"; copying preserves the file owner. After copying, "userl" has 15 capacity left returns "false"; "userl" does not have enough storage capacity left to perform copying operation returns "5"
returns "true"; this operation is done by "admin" with unlimited capacity returns ""; the file "/dir/file.small" already exists
returns ""; "userl" does not have enough storage capacity left to add this file returns ""; "user3" doesn't exist
returns "0"; all files owned by "userl" can fit into the new capacity of 300 bytes returns "2"; the files "/dir/file.big" and "/dir/another/file.med"
should be deleted so the remaining files owned by "userl" can fit into the new capacity of 50 bytes
returns ""; "user2" doesn't
Level 4:
Level 4
Implement support for file compression.
• COMPRESS_FILE <userId> ‹name> - Should compress the file name if it belongs to userid. The compressed file should be replaced with a new file named ‹name>. COMPRESSED. The size of the newly created file should be equal to the half of the original file. The size of all files is guaranteed to be even, so there should be no fractional calculations. It is also guaranteed that name for this operation never points to a compressed file - i.e., it never ends with COMPRESSED. Compressed files should be owned by userId - the owner of the original file. Returns a string representing the remaining storage capacity for userId if the file was compressed successfully or an empty string otherwise.
Note that because file names can only contain lowercase letters, compressed files cannot be added via ADD_FILE .
It is guaranteed that all COPY_FILE operations will preserve the suffix COMPRESSED
• DECOMPRESS _FILE ‹userId> ‹name> - Should revert the compression of the file name if it belongs to userid. It is guaranteed that name for this operation always ends with .COMPRESSED . If decompression results in the userid exceeding their storage capacity limit or a decompressed version of the file with the given name already exists, the operation fails. Returns a string
representing the remaining capacity of userid if the file was decompressed successfully or an empty string
compressing, decomporessing files
[
["ADD_USER",
["DECOMPRESS_FILE", "userl", "/file.mp4. COMPRESSED" ]
Queries
Explanations


"userl", "1000"],
returns "true"
["ADD_USER", "user2", "5000"),
returns "true"
['ADD_FILE_BX", "userl", "/dir/file.mp4", "500"],
returns "500"
['COMPRESS_FILE", "user2", "/dix/file.mp4"],
returns ""; the file "/dix/file.mp4" is owned by
['COMPRESS_FILE", "user3", "/dir/file.mp4"],
returns ""; "user3" doesn't exist
["COMPRESS_FILE", "userl",
"/foldex/non_existing_file"),
returns ""; the file "/folder/non_existing_file" doesn't exist
['COMPRESS_FILE", "userl", "/dir/file.mp4"],
returns "750"; the file "/dir/file.mp4" is compressed to size = 500 / 2 = 250 bytes
["GET _FILE_SIZE", "/dix/file.mp4. COMPRESSED"1,
returns "250"
["GET_FILE_SIZE", "/dix/file.mp4"],
returns ""
["COPY_FILE", "/dir/file.mp4.COMPRESSED", COMPRESSED", COMPRESSED"] COMPRESSED"],
returns "true"
["ADD_FILE_BY", "userl",
"/dir/file.mp4", "500"l,
returns "0"
["DECOMPRESS_FILE", "userl", "/dir/file.mp4.COMPRESSED"],
returns ""; "userl" does not have enough storage capacity to decompress the file
["UPDATE _CAPACITY", "useri", "2000"],
returns "0"
["DECOMPRESS_FILE", "user2", "/dir/file.mp4. COMPRESSED" 1,
returns ""; the file "/dir/file.mp4.COMPRESSED" is owned by "user1"
["DECOMPRESS _FILE", "user3", "/dir/file.mp4.COMPRESSED"1,
returns ""; "user3" doesn't exist
["DECOMPRESS _FILE", "userl", "/dir/file.mp4.COMPRESSED"1,

*/

type File struct {
	path              string
	owner             []string
	createdTime       string
	size              int
	version           string
	latestVersion     int
	otherVersionFiles []*File
}

type FileKey struct {
	path    string
	version string
}

type User struct {
	userId     string
	ownedFiles []string
	capacity   int
	usedSize   int
}
type CloudStorage struct {
	cs    map[FileKey]*File
	trash map[FileKey]*File
	users map[string]*User
}

func Init() *CloudStorage {
	return &CloudStorage{cs: make(map[FileKey]*File), trash: make(map[FileKey]*File),
		users: make(map[string]*User)}
}

func convertStrInt(num string) int {
	nFileSize, _ := strconv.Atoi(num)
	return nFileSize
}

func convertIntStr(num int) string {
	rt := fmt.Sprintf("%d", num)
	return rt
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (this *CloudStorage) Add(filePath string, size string) bool {
	k := FileKey{path: filePath}
	if _, ok := this.cs[k]; !ok {
		file := File{path: filePath, size: convertStrInt(size), latestVersion: 0}
		this.cs[k] = &file
		// this.cs[filePath]["size"] = size
		// this.cs[filePath]["path"] = filePath
		return true
	} else {

		if this.cs[k].otherVersionFiles == nil {
			this.cs[k].otherVersionFiles = make([]*File, 0)
		}
		this.cs[k].latestVersion++
		file := File{path: filePath, size: convertStrInt(size), latestVersion: this.cs[k].latestVersion, version: convertIntStr(this.cs[k].latestVersion)}

		this.cs[k].otherVersionFiles = append(this.cs[k].otherVersionFiles, &file)
		// this.cs[k].latestVersion = convertStrInt(this.cs[k].otherVersionFiles[len(this.cs[k].otherVersionFiles)-1].version)
		// version := this.cs[k].latestVersion + 1
		// this.cs[k].latestVersion++
		// file := File{path: filePath, size: convertStrInt(size), latestVersion: version}
		// this.cs[k] = &file

		return true
	}
}

func (this *CloudStorage) DeleteVersion(filePath string, version string) string {

	k := FileKey{path: filePath}

	if _, ok := this.cs[k]; !ok {
		return ""
	}

	if version == "0" {
		if len(this.cs[k].otherVersionFiles) > 0 {
			temp := &File{path: this.cs[k].otherVersionFiles[0].path, size: this.cs[k].otherVersionFiles[0].size, version: "0"}
			temp.otherVersionFiles = make([]*File, 0)

			for _, entry := range this.cs[k].otherVersionFiles[1:] {
				temp1 := &File{path: entry.path, size: entry.size, version: convertIntStr(convertStrInt(entry.version) - 1)}
				temp.otherVersionFiles = append(temp.otherVersionFiles, temp1)
			}

			delete(this.cs, k)

			this.cs[k] = temp
			this.cs[k].latestVersion = this.cs[k].otherVersionFiles[len(this.cs[k].otherVersionFiles)-1].latestVersion
		} else {
			delete(this.cs, k)
		}

	} else {

		for index, entry := range this.cs[k].otherVersionFiles {

			if entry.version == version {
				rtSize := entry.size
				this.cs[k].otherVersionFiles = append(this.cs[k].otherVersionFiles[:index], this.cs[k].otherVersionFiles[index+1:]...)

				for i := index; i < len(this.cs[k].otherVersionFiles); i++ {
					this.cs[k].otherVersionFiles[i].version = convertIntStr(convertStrInt(this.cs[k].otherVersionFiles[i].version) - 1)
				}
				this.cs[k].latestVersion = convertStrInt(this.cs[k].otherVersionFiles[len(this.cs[k].otherVersionFiles)-1].version)
				return convertIntStr(rtSize)

			}

		}

	}

	//  := this.cs[k].latestVersion + 1
	// this.cs[k].latestVersion++
	// file := File{path: filePath, size: convertStrInt(size), latestVersion: version}
	// this.cs[k] = &file

	return "false"

}

func (this *CloudStorage) TopN(filePrefix string, n string) bool {

	temp := make([]*File, 0)
	for k, _ := range this.cs {
		var f *File
		if strings.HasPrefix(k.path, filePrefix) {
			if this.cs[k].latestVersion == 0 {
				f = &File{path: this.cs[k].path, size: this.cs[k].size}
			} else {
				for _, val := range this.cs[k].otherVersionFiles {
					if convertStrInt(val.version) == this.cs[k].latestVersion {
						f = &File{path: val.path, size: val.size}
					}
				}

			}

			temp = append(temp, f)

		}
	}

	sort.Slice(temp, func(i, j int) bool {
		if temp[i].size == temp[j].size {
			return temp[i].path < temp[j].path
		}
		return temp[i].size > temp[j].size
	})

	return true

}

func (this *CloudStorage) Copy(sourcePath string, targetPath string) bool {
	tk := FileKey{path: targetPath}
	k := FileKey{path: sourcePath}
	if _, ok := this.cs[tk]; ok {
		return false
	}

	if _, ok := this.cs[k]; !ok {
		return false
	}

	contains := strings.Contains(k.path, ".")
	if !contains {
		return false
	}

	original := this.cs[k]
	this.cs[tk] = &File{path: targetPath, owner: append([]string(nil), original.owner...), size: original.size}
	if original.owner != nil {
		for _, user := range original.owner {
			this.users[user].ownedFiles = append(this.users[user].ownedFiles, targetPath)
			this.users[user].usedSize += this.cs[tk].size
		}
	}
	// this.cs[targetPath].path = targetPath

	return true

}

func (this *CloudStorage) GetFileSzie(sourcePath string) string {

	k := FileKey{path: sourcePath}
	if _, ok := this.cs[k]; !ok {
		return ""
	} else {
		return convertIntStr(this.cs[k].size)
	}
}

func (this *CloudStorage) FindFile(prefix string, surffix string) string {
	var fileList []*File
	for k, val := range this.cs {
		ifPrefix := strings.HasPrefix(k.path, prefix)
		ifSurfix := strings.HasSuffix(k.path, surffix)
		if ifPrefix && ifSurfix {
			// temp := make(map[string]interface{})
			// temp[k] = val
			fileList = append(fileList, val)
		}
	}

	sort.Slice(fileList, func(i, j int) bool {
		return fileList[i].size > fileList[j].size
	})

	strRet := ""

	for _, val := range fileList {
		s := convertIntStr(val.size)
		p := val.path
		strRet = strRet + p + "(" + s + ")" + " ,"
	}

	return strRet

}

func (this *CloudStorage) AddUsers(userId string, capacity string) bool {
	if _, ok := this.users[userId]; ok {
		return false
	} else {
		this.users[userId] = &User{userId: userId, capacity: convertStrInt(capacity), usedSize: 0}
		return true
	}
}

func (this *CloudStorage) AddFileBy(userId string, file string, filesize string) bool {

	k := FileKey{path: file}
	if _, ok := this.users[userId]; !ok {
		return false
	}

	var addresult bool
	if _, ok := this.cs[k]; !ok {
		addresult = this.Add(file, filesize)
	}

	if addresult == false {
		return false
	}

	// rem := this.users[userId].remaining - convertStrInt(filesize)

	usedSize := this.users[userId].usedSize + convertStrInt(filesize)

	if usedSize > this.users[userId].capacity {
		return false
	}
	this.cs[k].owner = append(this.cs[k].owner, userId) // don't need to put owner as part of files
	this.users[userId].ownedFiles = append(this.users[userId].ownedFiles, file)
	this.users[userId].usedSize = usedSize
	return true

}

func (this *CloudStorage) UpdateCapacity(userId string, capacity string) string {

	// if convertStrInt(capacity) <= this.users[userId].capacity {
	// 	return ""
	// }

	userFileList := this.users[userId].ownedFiles
	SortFiles := []*File{}

	for _, val := range userFileList {
		k := FileKey{path: val}
		SortFiles = append(SortFiles, this.cs[k])
	}

	sort.Slice(SortFiles, func(i, j int) bool {
		return SortFiles[i].size > SortFiles[j].size
	})

	tempTotal := this.users[userId].usedSize
	cutIndex := 0
	for index, file := range SortFiles {
		tempTotal = tempTotal - file.size
		if tempTotal <= convertStrInt(capacity) {
			cutIndex = index
			break
		}
	}

	this.users[userId].ownedFiles = []string{}

	newUsedSize := 0
	for _, val := range SortFiles[cutIndex+1:] {
		this.users[userId].ownedFiles = append(this.users[userId].ownedFiles, val.path)
		k := FileKey{path: val.path}
		newUsedSize += this.cs[k].size
	}

	this.users[userId].usedSize = newUsedSize
	this.users[userId].capacity = convertStrInt(capacity)

	return convertIntStr(cutIndex + 1)

}

func (this *CloudStorage) CompressFile(userId string, file string) string {

	//check if file exist, if user exist, if userId owns the file

	// newFileName := file + ".COMPRESSED"

	// this.cs[newFileName] = &File{path:newFileName, size: }
	k := FileKey{path: file}
	delete(this.cs, k)
	return ""

}

func (this *CloudStorage) DeCompressFile(userId string, file string) string {

	//check if file exist, if user exist, if userId owns the file

	// newFileName := file + ".COMPRESSED"

	// this.cs[newFileName] = &File{path:newFileName, size: }
	// output := strings.TrimSuffix(file, ".COMPRESSED")
	k := FileKey{path: file}

	if !containsString(this.cs[k].owner, userId) {
		return "false"
	}

	for _, owner := range this.cs[k].owner {
		if _, ok := this.users[owner]; ok {
			if this.cs[k].size+this.users[owner].usedSize > this.users[owner].capacity {
				return "false"
			}
		}
	}
	delete(this.cs, k)
	return ""

}
