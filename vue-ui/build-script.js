// const fs = require("fs")
import fs from 'node:fs'
// const { exec } = require('child_process');
import { exec } from 'node:child_process'
// console.log('only really used for local development, not dockerfile, it should already do it')

async function run() {
  await cli_exec()
  await post_script()
}

run()

async function cli_exec() {
  return new Promise((res, rej) => {
    exec('npm run build', (error, stdout, stderr) => {
      if (error) {
        console.error(`error: ${error.message}`)
        // return
        rej(error)
      }

      if (stderr) {
        console.error(`stderr: ${stderr}`)
        // return
        rej(stderr)
      }

      console.log(`stdout:\n${stdout}`)
      res(true)
    })
  })
}

async function post_script() {
  fs.cp('./dist/index.html', '../games-go/templates/index.html', (err) => {
    if (err) console.error(err)
  })
  fs.cp('./src/output.css', '../games-go/assets/output.css', (err) => {
    if (err) console.error(err)
  })
  fs.cp('./dist/assets/', '../games-go/assets/', { recursive: true }, (err) => {
    if (err) console.error(err)
  })
}
