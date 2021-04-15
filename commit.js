const exec = require('child_process').execSync

run("git add .")

let changedFiles = run("git status -s").split("\n").map(e => e.split(" ")[e.split(" ").length - 1])

let files = []

for (const changedFile of changedFiles) {
    const args = changedFile.split("/")
    let type = ""
    let message = ""
    if (args.length === 1) {
        const name = args[0]
        args[0] = "snippets"
        args.push(name)
    }

    const status = run(`git status ${changedFile} -s`).split("")[0]

    switch (status) {
        case "A":
            type = "feat"
            message = "added"
            break
        case "M":
            type = "refactor"
            message = "updated"
            break
        case "D":
            type = "refactor"
            message = "deleted"
            break
        case "R":
            type = "refactor"
            message = "renamed"
            break
        case "C":
            type = "refactor"
            message = "copied"
            m
        default:
            break;
    }

    files.push({
        path: changedFile,
        name: args[0] === "snippets" ? args[args.length - 1] : args[args.length - 2],
        language: args[0],
        commitType: type,
        commitMessage: message
    })
}

for (let file of files) {
    run(`git restore --staged ${file.path}`)
}

for (let file of files) {
    run(`git add ${file.path}`)
    console.log(run(`git commit -m "${file.language}: ${file.commitMessage} ${file.name}"`))
    // console.log(run(`git commit -m "${file.commitType}(${file.language}): ${file.commitMessage} ${file.name}"`))
}

function run(cmd) {
    console.log(`# Running command: "${cmd}"`)
    return exec(cmd, {encoding: 'utf-8'}).trim()
}
