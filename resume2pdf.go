package main

import (
"fmt"
"os"
"github.com/sopov/resumeio2pdf"
)

func main() {
secure_id := os.Args[1]
pdf_file, err := resumeio2pdf.Convert(secure_id)
if err != nil {
fmt.Println(err)
return
}

fmt.Println("Downloading resume...")
pdf_file.WriteTo(os.Stdout)
}
