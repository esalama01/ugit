# ugit

A lightweight, functional implementation of the Git source control system developed in Go. This project serves as a deep dive into the internal mechanics of version control, focusing on object storage, indexing, and the Merkle tree architecture.

## Core Architecture

The project implements the fundamental pillars of Git:

### 1. Content-Addressable Storage (CAS)
Everything in `ugit` is an object. Data is stored in the `.ugit/objects` directory, indexed by the SHA-1 hash of its content. 
- **Blobs**: Store file data with a zlib-compressed header.
- **Trees**: Represent directory structures, linking filenames to blob or tree hashes.
- **Commits**: Store snapshots of the root tree along with author metadata and parent history.

### 2. The Staging Area (Index)
The `index` file acts as the single source of truth between the working directory and the repository. It tracks file paths, SHA-1 hashes, and metadata to determine what has changed.

### 3. Merkle Tree Construction
During a commit, `ugit` performs a post-order traversal to build a hierarchy of trees. This ensures that a single root hash represents the entire state of the project at that specific moment.

## Features Implemented

- **Initialization**: `ugit init` sets up the internal directory structure (`objects`, `refs`, `heads`).
- **Object Hashing**: `ugit hash-object` generates unique identifiers for content with optional write support.
- **Staging**: `ugit add` updates the index and prepares snapshots.
- **Status Tracking**: `ugit status` compares the working directory, index, and latest commit to identify untracked, modified, or deleted files.
- **Snapshotting**: `ugit commit -m "message"` creates a permanent record of the staged changes.
- **Difference Engine**: `ugit diff` implements the Myers diff algorithm to visualize line-by-line changes between the index and the working directory.

## Getting Started

### Prerequisites
- Go 1.21 or higher installed on your system.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/esalama01/ugit.git ugit
   cd ugit
   ```

2. Build the executable:
   ```bash
   go build -o ugit
   ```

3. (Optional) Move the binary to your path:
   ```bash
   sudo mv ugit /usr/local/bin/
   ```

## Usage

Initialize a new repository:
```bash
./ugit init
```

Check the status:
```bash
./ugit status
```

Stage a file:
```bash
./ugit add <filename>
```

Create a commit:
```bash
./ugit commit -m "Your commit message"
```

View differences:
```bash
./ugit diff
```

Calculate the hash of a file:
```bash
./ugit hash-object <filename>
# Use -w to write the object to the database
./ugit hash-object -w <filename>
```

## Internal Directory Structure

```plaintext
.ugit/
├── HEAD          # Points to the current active branch
├── index         # Binary-like file representing the staging area
├── objects/      # Content-addressable database (sharded by first 2 chars of SHA-1)
└── refs/
    └── heads/    # Branch pointers (e.g., main)
```

## Implementation Details

- **Language**: Go 1.21+
- **Compression**: Standard zlib for object storage.
- **Hashing**: crypto/sha1 for generating 40-character hexadecimal identifiers.
- **Diff Logic**: Utilizes the Myers algorithm for efficient Longest Common Subsequence (LCS) detection.

## Future Work

Due to the significant time and effort spent mastering the core local architecture, network-related features like `ugit push` and `ugit clone` have not been implemented in this version. However, expanding the tool to support the Git Transfer Protocol and Smart HTTP handshake remains a primary goal for future updates.
