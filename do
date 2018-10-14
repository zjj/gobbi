#!/bin/bash
set -e
set -u

exampleDir=internal/example
tags="glib_2.48 gdk_3.4 gio_2.32 gtk_3.18"

usage () {
    echo "Usage:"
    echo "  do -h                       Show this help"
    echo "  do build                    Generate libraries, and run tests"
    echo "  do generate                 Generate libraries"
    echo "  do test                     Run tests"
    echo "  do examples                 List the available examples"
    echo "  do example <example-name>   Run an example"

    exit
}

example () {
    exampleName=$1
    go run -tags "$tags" $exampleDir/$exampleName/*.go
}

build () {
    generate && \
    echo "" && \
    test
}

generate () {
    go run internal/cmd/generate/*.go
}

test () {
    go test -tags "$tags" github.com/pekim/gobbi/internal/test/...
}

examples () {
  for example in $(find $exampleDir -name main.go -printf '%h\n') ; do
    basename $example
  done
}

processOptions () {
  while getopts ":h" opt; do
    case ${opt} in
      h )
        usage
        ;;
      \? )
        echo "Invalid Option: -$OPTARG" 1>&2
        echo ""
        usage
        ;;
    esac
  done
  shift $((OPTIND -1))

  if [ "$#" -lt 1 ]; then
    usage
  fi
}

processSubcommand () {
  subCommand=$1
  shift

  case "$subCommand" in
    build)
      build
      ;;
    generate)
      generate
      ;;
    test)
      test
      ;;
    examples)
      examples
      ;;
    example)
      example $@
      ;;
    *)
      echo "Invalid subcommand: $subCommand" 1>&2
      echo ""
      usage
      ;;
  esac
}

processOptions "$@"
processSubcommand "$@"
