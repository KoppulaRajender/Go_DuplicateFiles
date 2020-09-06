# Go_DuplicateFiles
Go language code for checking duplicate files in the provided directory

#Docker Commands

docker pull koppula18/go:duplicatefilescheck-v1.0

docker run -it --mount src=<duplicatefiles_directory>,target=/go/mydir,type=bind koppula18/go:duplicatefilescheck-v1.0 /go/mydir
