var findWords = function (board, words) {
    let result = [];
    let loop = true;
    for (let str of words) {
        loop = true;
        for (let i = 0; i < board.length; i++) {
            for (let j = 0; j < board[0].length; j++) {
                if (loop && board[i][j] === str[0] && helper(JSON.parse(JSON.stringify(board)), str, i, j, 0)) {
                    result.push(str);
                    loop = false;
                }
            }
        }
    }
    return result.filter((e,i,a) => a.indexOf(e) == i);
};

function helper(board, word, i, j, index) {
    if (i < 0 || j < 0 || i >= board.length || j >= board[0].length || index >= word.length) {
        return false;
    } else if (board[i][j] != word[index]) {
        return false;
    } else {
        if (index == word.length - 1) {
            return true;
        }
        let temp = board[i][j];
        board[i][j] = '#';
        let res = helper(board, word, i + 1, j, index + 1) || helper(board, word, i - 1, j, index + 1) ||
            helper(board, word, i, j + 1, index + 1) || helper(board, word, i, j - 1, index + 1);
        board[i][j] = temp;
        return res;
    }
}