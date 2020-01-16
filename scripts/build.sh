#!/usr/bin/env bash

function showUsageAndExit() {
    echo "Insufficient or invalid options provided"
    echo
    echo "Usage: "$'\e[1m'"./build.sh -t [target-file] -v [build-version] -f"$'\e[0m'
    echo -en "  -t\t"
    echo "[REQUIRED] Target file to build."
    echo -en "  -v\t"
    echo "[REQUIRED] Build version. If not specified a default value will be used."
    echo -en "  -f\t"
    echo "[OPTIONAL] Cross compile for all the list of platforms. If not specified, the specified target file will be cross compiled only for the autodetected native platform."

    echo
    echo "Ex: "$'\e[1m'"./build.sh -t wum.go -v 1.0.0 -f"$'\e[0m'" - Builds WUM Client for version 1.0.0 for all the platforms"
    echo
    exit 1
}

while getopts :t:v:f FLAG; do
  case ${FLAG} in
    t)
      target=$OPTARG
      ;;
    v)
      build_version=$OPTARG
      ;;
    f)
      full_build="true"
      ;;
    \?)
      showUsageAndExit
      ;;
  esac
done

sourcePath=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
commandPath=$(cd "$(dirname ${sourcePath})" && pwd)
rootPath=$(cd "$(dirname ${sourcePath})" && pwd)

if [[ ! -e "${rootPath}/cmd/${target}" ]]; then
  echo "Target file is needed. "
  showUsageAndExit
  exit 1
fi

if [[ -z "${build_version}" ]]
then
  echo "Build version is needed. "
  showUsageAndExit
fi


buildDir="build/target"
buildPath="${rootPath}/${buildDir}"

echo "Cleaning build path ${buildDir}..."
rm -rf ${buildPath}

if [[ "${full_build}" == "true" ]]; then
    platforms="darwin/amd64/macosx/x64 linux/amd64/linux/x64 windows/amd64/windows/x64"
else
    platform=$(uname -s)
    if [[ "${platform}" == "Linux" ]]; then
        platforms="linux/amd64/linux/x64"
    elif [[ "${platform}" == "Darwin" ]]; then
        platforms="darwin/amd64/macosx/x64"
    else
        platforms="windows/amd64/windows/x64"
    fi
fi

for platform in ${platforms}
do
    split=(${platform//\// })
    goos=${split[0]}
    goarch=${split[1]}
    pos=${split[2]}
    parch=${split[3]}

    output="${target}_${goos}"

    # add exe to windows output
    [[ "windows" == "${goos}" ]] && output="${output}.exe"

    echo -en "\t - ${goos}/${goarch}..."

    zipfile="${target}-${build_version}-${pos}-${parch}"
    zipdir="${buildPath}/${target}"
    mkdir -p ${zipdir}

    # set destination path for binary
    destination="$zipdir/bin/${output}"

    #echo "GOOS=$goos GOARCH=$goarch go build -x -o $destination $target"
    GOOS=${goos} GOARCH=${goarch} go build -gcflags=-trimpath=${GOPATH} -asmflags=-trimpath=${GOPATH} \
    -ldflags "-X main.deploymentChecker=$build_version -X 'main.buildDate=$(date -u '+%Y-%m-%d %H:%M:%S UTC')'" \
    -o ${destination} ${rootPath}/cmd/${target}/main.go


    pwd=`pwd`
    cd ${buildPath}
    tar czf "$zipfile.tar.gz" ${target} > /dev/null 2>&1
    rm -rf ${target}
    cd ${pwd}
    echo -en $'\e[1m\u2714\e[0m'
    echo
done

echo "Build complete!"
