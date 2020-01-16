const fx = require('mkdir-recursive');
const copydir = require('copy-dir');
const path = require('path');

const goroot = process.env.GOROOT;
const workspaceDir = process.env.GITHUB_WORKSPACE;

var projectSource = path.join(goroot, "src", "github.com", "wso2", "information-extractor");
fx.mkdir(projectSource, function (err) {
    console.log("Created project source " + projectSource);

    copydir.sync(workspaceDir, projectSource, {
        utimes: true,
        mode: true,
        cover: true
    });
    console.log("Copied project to $GOROOT.")
});
