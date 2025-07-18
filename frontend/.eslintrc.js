module.exports = {
    root: true,
    env: {
      node: true,
      browser: true
    },
    extends: [
      'plugin:vue/vue3-essential',
      'eslint:recommended'
    ],
    parserOptions: {
      ecmaVersion: 2020
    },
    rules: {
      'vue/multi-word-component-names': 'off'
    }
  };
  