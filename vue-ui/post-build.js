// const fs = require("fs")
import fs from 'node:fs'
console.log('only really used for local development, not dockerfile, it should already do it')

async function post_script() {
  fs.cp('./dist/index.html', '../games-go/templates/index.html', (err) => {
    console.error(err)
  })
  fs.cp('./src/output.css', '../games-go/assets/output.css', (err) => {
    console.error(err)
  })
  fs.cp('./dist/assets/', '../games-go/assets/', { recursive: true }, (err) => {
    console.error(err)
  })
}

post_script()
