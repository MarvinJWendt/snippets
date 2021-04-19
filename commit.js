const exec = require('child_process').execSync

const flags = process.argv.slice(2);

run("go run ./ci/main.go", false)
run("git add -u .")

let changedFiles = run("git status -s").split("\n").map(e => e.split(" ")[e.split(" ").length - 1])

let files = []

for (const changedFile of changedFiles) {
    const args = changedFile.split("/")
    let message = ""
    if (args.length === 1) {
        const name = args[0]
        args[0] = "snippets"
        args.push(name)
    }

    const status = run(`git status ${changedFile} -s`).split("")[0]

    switch (status) {
        case "A":
            message = "added"
            break
        case "M":
            message = "updated"
            break
        case "D":
            message = "deleted"
            break
        case "R":
            message = "renamed"
            break
        case "C":
            message = "copied"
            break
        case "?":
            message = "added"
            break
        default:
            break
    }

    files.push({
        path: changedFile,
        scriptName: args[0] === "snippets" ? args[args.length - 1] : args[args.length - 2],
        name: args[args.length - 1],
        language: args[0],
        commitMessage: message
    })
}

for (let file of files) {
    if (!file.path.endsWith("/")) run(`git restore --staged ${file.path}`)
}

for (let file of files) {
    run(`git add ${file.path}`)
    console.log(run(`git commit -m "${file.language}: ${file.commitMessage} ${file.name} in ${file.language === "snippets" ? "snippets" : file.scriptName}"`))
}

if (flags[0] === "push") console.log(run("git push"))

function run(cmd, quite) {
    if (quite !== true) {
        console.log(`# Running command: "${cmd}"`)
    }
    return exec(cmd, {encoding: 'utf-8'}).trim()
}
