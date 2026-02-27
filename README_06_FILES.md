# 📁 README_06 — File I/O in Go

> **File Reference:** `files/file.go`, `files/a.txt`

---

## 1. The `os` Package

Go's `os` package is the primary way to interact with the file system.
It wraps OS-level syscalls and works on Windows, Linux, and macOS.

```go
import "os"
```

### Most Used Functions

| Function | Purpose |
|----------|---------|
| `os.Open(name)` | Open file for **reading** only |
| `os.Create(name)` | Create/truncate file for **writing** |
| `os.OpenFile(name, flag, perm)` | Full control (read, write, append) |
| `os.ReadFile(name)` | Read **entire** file into `[]byte` |
| `os.WriteFile(name, data, perm)` | Write **entire** file from `[]byte` |
| `os.Remove(name)` | Delete a file |
| `os.Mkdir(name, perm)` | Create a directory |
| `os.MkdirAll(path, perm)` | Create nested directories |
| `os.Stat(name)` | Get file/dir info |

---

## 2. Opening & Reading Files

### Method A — `os.ReadFile` (Simplest — read all at once)

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("a.txt")
    if err != nil {
        panic(err) // or log.Fatal(err)
    }
    fmt.Println(string(data)) // convert []byte to string
}
```

### Method B — `os.Open` + `f.Read` (Manual buffer)

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Open("a.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close() // ALWAYS defer Close!

    buf := make([]byte, 10) // read 10 bytes at a time
    n, err := f.Read(buf)
    if err != nil {
        panic(err)
    }
    fmt.Println("Bytes read:", n)
    fmt.Println("Content:", string(buf[:n]))
}
```

### Method C — `bufio.Scanner` (Read line by line — most common)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    f, err := os.Open("a.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    lineNum := 1
    for scanner.Scan() {
        fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
        lineNum++
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
```

---

## 3. Getting File Info (`os.Stat`)

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fileInfo, err := os.Stat("a.txt")
    if err != nil {
        panic(err)
    }

    fmt.Println("Name:    ", fileInfo.Name())    // a.txt
    fmt.Println("Size:    ", fileInfo.Size())    // size in bytes
    fmt.Println("IsDir:   ", fileInfo.IsDir())   // false
    fmt.Println("ModTime: ", fileInfo.ModTime()) // last modified
    fmt.Println("Mode:    ", fileInfo.Mode())    // file permissions
}
```

### Check if File Exists

```go
func fileExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}
```

---

## 4. Writing Files

### Method A — `os.WriteFile` (Simple — write all at once)

```go
package main

import "os"

func main() {
    data := []byte("Hello, Go!\nThis is line 2.\n")
    err := os.WriteFile("output.txt", data, 0644)
    if err != nil {
        panic(err)
    }
    // File created with content
}
```

**Permission `0644`:**
- Owner: read + write
- Group: read only
- Others: read only

### Method B — `os.Create` + `f.WriteString`

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Create("output.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    f.WriteString("Hello, Go!\n")
    f.WriteString("Second line\n")

    bytes := []byte("Written as bytes\n")
    n, err := f.Write(bytes)
    fmt.Println("Wrote", n, "bytes")
}
```

### Method C — `bufio.Writer` (Buffered — efficient for large writes)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    f, err := os.Create("output.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    writer := bufio.NewWriter(f)
    fmt.Fprintln(writer, "Buffered line 1")
    fmt.Fprintln(writer, "Buffered line 2")
    writer.Flush() // MUST flush to write buffer to disk!
}
```

---

## 5. `os.OpenFile` — Full Control

`os.OpenFile` lets you open with specific **flags**:

```go
f, err := os.OpenFile("log.txt",
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, // flags
    0644, // permissions
)
```

### Common Flags

| Flag | Meaning |
|------|---------|
| `os.O_RDONLY` | Read only |
| `os.O_WRONLY` | Write only |
| `os.O_RDWR` | Read and write |
| `os.O_CREATE` | Create if not exists |
| `os.O_APPEND` | Append to end |
| `os.O_TRUNC` | Truncate to zero on open |

### Append to a File

```go
f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    panic(err)
}
defer f.Close()

f.WriteString("New log entry\n")
```

---

## 6. Working with Directories

### List Directory Contents

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    dir, err := os.Open("../")
    if err != nil {
        panic(err)
    }
    defer dir.Close()

    // Read up to 5 entries (-1 for all)
    entries, err := dir.ReadDir(-1)
    if err != nil {
        panic(err)
    }

    for _, entry := range entries {
        if entry.IsDir() {
            fmt.Println("[DIR] ", entry.Name())
        } else {
            fmt.Println("[FILE]", entry.Name())
        }
    }
}
```

### Create and Remove Directories

```go
// Create a single directory
os.Mkdir("logs", 0755)

// Create nested directories (like mkdir -p)
os.MkdirAll("logs/2024/january", 0755)

// Remove an empty directory
os.Remove("logs")

// Remove directory and all its contents
os.RemoveAll("logs")
```

---

## 7. Deleting Files (from your `file.go`)

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.Remove("a2.txt")
    if err != nil {
        panic(err)
    }
    fmt.Println("file deleted successfully")
}
```

---

## 8. Copying a File (bufio pattern)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func copyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destFile.Close()

    reader := bufio.NewReader(sourceFile)
    writer := bufio.NewWriter(destFile)

    for {
        b, err := reader.ReadByte()
        if err != nil {
            break // EOF or error
        }
        writer.WriteByte(b)
    }

    return writer.Flush()
}

func main() {
    err := copyFile("a.txt", "a_copy.txt")
    if err != nil {
        panic(err)
    }
    fmt.Println("Copied successfully")
}
```

> **Better alternative:** Use `io.Copy(dst, src)` — it does this automatically!

```go
import "io"

src, _ := os.Open("a.txt")
defer src.Close()
dst, _ := os.Create("a_copy.txt")
defer dst.Close()

io.Copy(dst, src)
```

---

## ✅ Solved Examples

### Example 1 — Word Counter

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func countWords(filename string) (int, error) {
    f, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    defer f.Close()

    count := 0
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        words := strings.Fields(scanner.Text())
        count += len(words)
    }
    return count, scanner.Err()
}

func main() {
    n, err := countWords("a.txt")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Word count: %d\n", n)
}
```

---

### Example 2 — Simple Logger

```go
package main

import (
    "fmt"
    "os"
    "time"
)

type Logger struct {
    file *os.File
}

func NewLogger(filename string) (*Logger, error) {
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, err
    }
    return &Logger{file: f}, nil
}

func (l *Logger) Log(msg string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    line := fmt.Sprintf("[%s] %s\n", timestamp, msg)
    l.file.WriteString(line)
}

func (l *Logger) Close() {
    l.file.Close()
}

func main() {
    log, err := NewLogger("app.log")
    if err != nil {
        panic(err)
    }
    defer log.Close()

    log.Log("Application started")
    log.Log("User logged in")
    log.Log("Application stopped")

    fmt.Println("Logs written to app.log")
}
```

---

### Example 3 — Read Config File (Key=Value format)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func readConfig(filename string) (map[string]string, error) {
    config := make(map[string]string)
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "#") {
            continue // skip blank lines and comments
        }
        parts := strings.SplitN(line, "=", 2)
        if len(parts) == 2 {
            config[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
        }
    }
    return config, scanner.Err()
}

func main() {
    // config.txt:
    //   DB_HOST=localhost
    //   DB_PORT=5432
    //   APP_NAME=MyApp
    config, err := readConfig("config.txt")
    if err != nil {
        panic(err)
    }
    fmt.Println("DB_HOST:", config["DB_HOST"])
    fmt.Println("APP_NAME:", config["APP_NAME"])
}
```

---

## 🏋️ Practice Problems

### Problem 1 — Create, Write, Read, Delete
Write a program that:
1. Creates a file called `test.txt`
2. Writes 3 lines to it
3. Reads it back and prints line by line
4. Deletes the file

<details>
<summary>💡 Solution</summary>

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // 1. Create and Write
    f, err := os.Create("test.txt")
    if err != nil { panic(err) }
    f.WriteString("Line 1: Go is awesome\n")
    f.WriteString("Line 2: File I/O is easy\n")
    f.WriteString("Line 3: Practice makes perfect\n")
    f.Close()

    // 2. Read line by line
    f, err = os.Open("test.txt")
    if err != nil { panic(err) }
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    f.Close()

    // 3. Delete
    os.Remove("test.txt")
    fmt.Println("Done!")
}
```
</details>

---

### Problem 2 — Line Counter
Write a function `countLines(filename string) (int, error)` that counts the number of lines in a file.

<details>
<summary>💡 Solution</summary>

```go
func countLines(filename string) (int, error) {
    f, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    defer f.Close()

    count := 0
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        count++
    }
    return count, scanner.Err()
}
```
</details>

---

### Problem 3 — Append Timestamp Log
Write a program that appends a new timestamped entry to `log.txt` every time it runs (file not truncated).

<details>
<summary>💡 Solution</summary>

```go
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil { panic(err) }
    defer f.Close()

    entry := fmt.Sprintf("[%s] App ran\n", time.Now().Format(time.RFC3339))
    f.WriteString(entry)
    fmt.Println("Logged:", entry)
}
```
</details>

---

## 🔑 Key Takeaways

| Concept | Remember |
|---------|----------|
| Always `defer f.Close()` | Prevents resource leaks |
| `os.ReadFile` | Simplest for small files |
| `bufio.Scanner` | Best for line-by-line reading |
| `bufio.Writer` | Best for many small writes |
| `os.OpenFile` flags | Use `O_APPEND` to not overwrite |
| `os.WriteFile` perm | `0644` = owner rw, others r |
| `writer.Flush()` | MUST call or data stays in buffer |

---

*Next → [🔒 README_07_MUTEX.md](./README_07_MUTEX.md) — Mutex & Concurrency Safety*
