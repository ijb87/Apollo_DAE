const { VueLoaderPlugin } = require('vue-loader')
const nodeExternals = require('webpack-node-externals')


module.exports = {
  module: {
    rules: [
      {
        test: /\.vue$/,
        use: 'vue-loader'
      }
    ]
  },
  plugins: [
    new VueLoaderPlugin()
  ]
}