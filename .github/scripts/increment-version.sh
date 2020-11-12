#!/usr/bin/env bash

GITHUB_TOKEN=$1

# -------------------------------------- Environment Variables --------------------------------------
# GITHUB_ACTOR - The name of the person or app that initiated the workflow
# GITHUB_REPOSITORY - The owner and repository name
# -------------------------------------- Environment Variables --------------------------------------

#get highest tag number
GIT_VERSION=`git describe --abbrev=0 --tags`
PROJECT_VERSION=$(make version)

function getNewVersion() {
  VERSION_BITS=(${GIT_VERSION//./ })

  MAJOR=${VERSION_BITS[0]}
  MINOR=${VERSION_BITS[1]}
  PATCH=${VERSION_BITS[2]}
  PATCH=$((PATCH+1))

  NEW_VERSION="${MAJOR}.${MINOR}.${PATCH}"
  echo ${NEW_VERSION}
}

function pushChanges() {
  [[ -z "${GITHUB_TOKEN}" ]] && {
    echo 'Missing input "github_token: ${{ secrets.GITHUB_TOKEN }}".';
    exit 1;
  };
  NEW_VERSION=${1}

  git add Makefile

  git config --local user.email "actions@github.com"
  git config --local user.name "GitHub Actions"
  git commit -m "Create new release ${NEW_VERSION}"

  remote_repository="https://${GITHUB_ACTOR}:${GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}.git"
  git push "${remote_repository}" HEAD:master
}

if [[ ${PROJECT_VERSION} == ${GIT_VERSION} ]]; then
  NEW_VERSION=$(getNewVersion)
  cat Makefile | sed s/"BUILD_VERSION := ${PROJECT_VERSION}"/"BUILD_VERSION := ${NEW_VERSION}"/ > tmp.txt
  mv tmp.txt Makefile

  pushChanges ${NEW_VERSION}
fi

