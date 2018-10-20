const path = require('path');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

module.exports = {
  entry: './web/public/js/main.js',
  output: {
    filename: 'main.js',
    path: path.resolve(__dirname, 'web/dist')
  },
  module: {
    rules: [
        {
            test: /\.css$/,
            use: [
              {
                loader: MiniCssExtractPlugin.loader
              },
              "css-loader"
            ]
          }
    ]
  },
  plugins: [
    new MiniCssExtractPlugin({
        filename: "[name].css",
        chunkFilename: "[id].css"
      })
  ]
};