let fx = require('mkdir-recursive');
const copydir = require('copy-dir');

const goroot = process.env.GOROOT;
const workspaceDir = process.env.GITHUB_WORKSPACE;

let projectSource = goroot + "/src/github.com/wso2";
fx.mkdir(projectSource, function (err) {
    console.log("Created project source " + projectSource);

    copydir.sync(workspaceDir, projectSource, {
        utimes: true,
        mode: true,
        cover: true
    });
    console.log("Copied project to $GOROOT.")
});