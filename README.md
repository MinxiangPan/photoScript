# script for helping photo relate process

### match.go:
##### args:
| args | description | default | example |
| --- | --- | --- | --- |
| 1 | input path (full path) | N/A (error) | ~/dev/photos |
| 2 | extension for target file (used as match checker) | .hif | .hif |
| 3 | extension for process file (used for delete any unmatch file) | .arw | .arw |
##### run (example)
```
go run match.go ~/dev/photos
go run match.go ~/dev/photos ".heif" ".raw"
```
