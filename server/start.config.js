module.exports = {
  apps: [
    {
      name: 'server',
      script: './app.js',
      exec_mode: 'cluster',
      watch: false,
      env: {
        NODE_ENV: 'production',
        REDIS_HOST: 'redis',
        REDIS_PORT: 6379,
        MYSQL_HOST: 'mysql',
        MYSQL_PORT: 3306,
        MYSQL_USER: 'user',
        MYSQL_PASSWORD: '123456a',
        MYSQL_DATABASE: 'data_analysis_system',
      }
    }
  ]
}