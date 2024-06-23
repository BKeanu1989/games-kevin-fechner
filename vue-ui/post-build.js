// const fs = require("fs")
import fs from 'node:fs'
console.log('post build')

async function post_script() {
  fs.cp('./dist/index.html', '../games-go-blueprint/templates/index.html', (err) => {
    console.error(err)
  })
  fs.cp('./src/output.css', '../games-go-blueprint/assets/output.css', (err) => {
    console.error(err)
  })
  fs.cp('./dist/assets/', '../games-go-blueprint/assets/', { recursive: true }, (err) => {
    console.error(err)
  })
}

post_script()
