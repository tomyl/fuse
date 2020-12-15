package fuse

// Maximum file write size we are prepared to receive from the kernel.
//
// Linux 5.4.0 caps this value at 1024kB (FUSE_MAX_MAX_PAGES=256, 4kB pages).
// Default kernel value is 128kB (FUSE_DEFAULT_MAX_PAGES_PER_REQ=32, 4kB
// pages).
const maxWrite = 1024 * 1024
