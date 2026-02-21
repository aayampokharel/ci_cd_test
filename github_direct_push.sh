# alias push='./github_direct_push.sh'
git add . 
branchName=${2:-"testing-previous-concepts"}
git commit -m "${1:-"fix(yml):CI practice"}"
git push origin $branchName