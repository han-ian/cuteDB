# Home 
This project is forked from CuteDB, which is a simple BTree implemention.

Neondb is a compute-storage-seperating relation database base on postgres. It save BTree page on LSM like key-value store and the LSM database is store on S3.

Inspired by neaon, the block storage of Btree is abstracted and stored on kv. This is what I do in this repository.

### Limitations.
* Don't support scan.
* No concurrency control.
* No MVCC on neither key-value nor page blocks.
* No page block 
* No garbage collection on page block.

