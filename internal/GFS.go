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
