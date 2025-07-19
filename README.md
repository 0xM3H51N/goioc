# goioc

A CLI tool written in Go to extract Indicators of Compromise (IOCs) such as IP addresses, domains, URLs, and file hashes from text files or logs.

## 🔧 Project Status

🚧 **In development** — not stable yet. Actively being built as part of a portfolio project.

## 📁 Project Structure

<pre>
goioc/
├── cmd/         # CLI entry point
├── internal/    # Core logic (IOC extraction)
├── testdata/    # Sample files for testing
├── go.mod
├── go.sum
└── README.md
</pre>


## 🎯 Planned Features

- [x] Basic CLI structure with flags
- [ ] Extract IOCs from text files
- [ ] Support for multiple input formats (log, plain text, etc.)
- [ ] JSON output support
- [ ] Integration with threat intelligence APIs (optional)

## 🧠 Learning Goals

This project helps reinforce:

- Go CLI development
- File parsing and regex
- Writing modular Go code
- Git/GitHub project hygiene

## 📦 Getting Started

```bash
git clone https://github.com/0xM3H51N/goioc.git
cd goioc
go run ./cmd


