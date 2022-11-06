var floodFill = function (image, sr, sc, newColor) {
    helper(image, sr, sc, newColor, image[sr][sc]);
    return image;
};

function helper(image, sr, sc, newColor, oldColor) {

    if (sr < 0 || sc < 0 || sr >= image.length || sc >= image[0].length) {
        return
    }
    if (image[sr][sc] !== oldColor) {
        return
    }
    image[sr][sc] = newColor
    return helper(image, sr + 1, sc, newColor, oldColor) || helper(image, sr - 1, sc, newColor, oldColor)
        || helper(image, sr, sc + 1, newColor, oldColor) || helper(image, sr, sc - 1, newColor, oldColor);

}

console.log(floodFill([[1, 1, 1], [1, 1, 0], [1, 0, 1]], 1, 1, 2))

console.log(floodFill([[0, 0, 0], [0, 0, 0]], 0, 0, 2))