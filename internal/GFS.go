package internal

/*
Design a large cluster of File System. e.g, S3, Azure Blob, GFS, Dropbox
Product requirements:
User is able to upload the data to the file system
User is able to download data from the file system
The FS is able to scale horizontally.


Considerations:
1. Hosting Services, public or private?
2. Load Balance
3. Can we use S3/blob storage as backing up FS? leverate exisint cloud
4. Max storage size of servers, can it be expanded
5. How many servers, server/HW specs, any white box?
6. Replication, fail-over
7. How to handle file writes, concurrent?
8. Interface
9. ACL
10. Caching? needed or use linux file system

Appraoch
1. Divide files into small chunks and store to different servers.
2. Define chunk size, large or small


Say I have a 100M file, chunk size is 50M, so I will have 2 chunks.
Each data has 2 parts, preamble-metadata/payload.
metadata
{"chunkId", "timestamp","datasize","crc_checksum"}


*/

/*
 */
// const (
// 	MASTER_SERVICE = iota
// 	DATA_SERVICE   = iota
// )

// /*

// /foo/folder/file1.txt -> [chunk1][chunk2][chunk3]

// */

// type GFS struct {
// 	role        int
// 	data_sum    int
// 	chunk_index []string
// 	chunk_table map[string][]string
// }

// /*
// HearBeat message from chunk server to master
// Storage usage
// IP address
// Mac address
// Role
// */

// type ChunkServerStats struct {
// 	ip_addr       string
// 	mac_address   string
// 	disk_usage    float32
// 	network_usage float32
// }

// /*
// The method to return the chunkServer
// The method is master's job
// Input: ChunkID
// Output: Chunk Server IP address

// e.g, /foo/george/file1, 1x877d

// */

// func (this *GFS) GetDataServer(fileName string, chunkId uint64) string {
// 	/*
// 		the method will use some algorithm to find out the most
// 		avaiable chunk server, based on a few factors, such as
// 		disk usage, network usage, can use priority q
// 	*/
// 	this.chunk_table[fileName][chunkId] = "server1"
// 	return "server"
// }

// /*
// Read a file by using namespace, e.g, /foot/file1
// to get the chunk table, and then get all the chunks
// */
// func (this *GFS) GetFile()

// /*
// Read a file by using namespace, e.g, /foot/file1
// to get the chunk table, and then get all the chunks
// */

// func (this *GFS) PutFile()
