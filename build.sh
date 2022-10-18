#!/bin/bash

git add --all
git commit -m 'added new structs'
git push

git tag v1.0.20
git push --tags