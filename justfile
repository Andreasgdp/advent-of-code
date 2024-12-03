gen LANGUAGE YEAR DAY:
  #!/bin/bash
  [[ -s "$GVM_ROOT/scripts/gvm" ]] && source "$GVM_ROOT/scripts/gvm"
  gvm use go1.23.3
  mkdir -p {{YEAR}}/day{{DAY}}-{{LANGUAGE}}
  cp -R templates/{{LANGUAGE}}/* {{YEAR}}/day{{DAY}}-{{LANGUAGE}}
  curl --cookie "session=$ADVENT_OF_CODE_COOKIE" https://adventofcode.com/{{YEAR}}/day/{{DAY}}/input -o {{YEAR}}/day{{DAY}}-{{LANGUAGE}}/input.txt
  perl -i -pe 'chomp if eof' {{YEAR}}/day{{DAY}}-{{LANGUAGE}}/input.txt
  printf '* [Day {{DAY}}](https://adventofcode.com/{{YEAR}}/day/{{DAY}}): [Go]({{YEAR}}/day{{DAY}}-go/main.go)\n' | sed '55r /dev/stdin' README.md > README.tmp && mv README.tmp README.md

  cd {{YEAR}}/day{{DAY}}-{{LANGUAGE}}
  go mod init day{{YEAR}}-{{DAY}}
  go mod tidy
  go get github.com/stretchr/testify
  go get github.com/teivah/advent-of-code
  go get golang.org/x/exp

  cd ../..

  nvim {{YEAR}}/day{{DAY}}-{{LANGUAGE}}
