var webpack = require('webpack');

var BUILD_DIR = __dirname + '/build';
var APP_DIR = __dirname + '/src';
console.log("BUILD_DIR: " + BUILD_DIR);
console.log("APP_DIR: " + APP_DIR);

var config = {
  entry: APP_DIR + '/index.js',
  output: {
    path: BUILD_DIR,
    filename: 'static/js/main.js',
    publicPath: '/',
  },
  resolve: {
    extensions: ['.js', '.jsx', '.json']
  }
};

module.exports = config;
