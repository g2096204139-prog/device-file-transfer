package storage

// Cloud storage expansion notes:
//
// v1.0 uses LocalStorage, but the Storage interface is intentionally designed
// so future implementations can support object storage such as S3, Google Cloud
// Storage, Azure Blob Storage, or MinIO.
//
// A future cloud implementation should preserve these operations:
// - Save
// - List
// - Open
// - Delete
//
// For v2.0 resumable uploads, storage may need extra methods for multipart
// upload sessions, chunk persistence, and final object composition.
