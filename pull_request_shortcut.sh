pr() {
    # Check if a branch name argument is provided
    if [ "$#" -ne 1 ]; then
        echo "Usage: pr <target_branch_name>"
        return 1
    fi
    TARGET_BRANCH=$1

    # Get the current Git repository URL and format it
    REPO_URL=$(git config --get remote.origin.url)
    if [[ "$REPO_URL" == git@github.com:* ]]; then
        # SSH format
        REPO_URL=${REPO_URL#*:}
        REPO_URL=${REPO_URL%.git}
    elif [[ "$REPO_URL" == https://github.com/* ]]; then
        # HTTPS format
        REPO_URL=${REPO_URL#https://github.com/}
        REPO_URL=${REPO_URL%.git}
    else
        echo "Unsupported repository URL format: $REPO_URL"
        return 1
    fi

    # Get the current branch name
    CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)

    # URL for creating a new pull request
    PR_URL="https://github.com/$REPO_URL/compare/$TARGET_BRANCH...$CURRENT_BRANCH"

    # Open the pull request URL in the default browser
    open "$PR_URL"
}
