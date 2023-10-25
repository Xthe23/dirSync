# dirSync

This Go script provides a utility to copy the contents of one folder to another. If the destination folder contains files or subdirectories with the same name as those in the source folder, they will be replaced. However, any files or directories in the destination folder that do not have corresponding items in the source folder will remain untouched. In essence, the script merges the contents of the source folder into the destination folder, overwriting any overlapping files, but leaving non-overlapping files and directories in the destination folder intact.

## Features
- Recursively copies all files and subdirectories from the source folder to the destination folder.
- Overwrites files in the destination folder if they have the same name as those in the source folder.
- Preserves the file permissions from the source folder.
- Non-overlapping files and directories in the destination folder remain untouched.

## Usage

1. Set the `folderA` (source) and `folderB` (destination) variables in the script to the paths of your desired source and destination folders.
2. Run the script using the command:
   ```bash
   go run <script-name>.go

