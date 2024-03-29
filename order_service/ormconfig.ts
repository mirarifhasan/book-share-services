require('dotenv').config({
  path: process.env.ENV_PATH,
});
console.log('driver ', process.env.DB_DRIVER);

const ormConfig: any = {
  type: process.env.DB_DRIVER,
  host: process.env.DB_HOST,
  port: process.env.DB_PORT,
  username: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
  entities: ['dist/**/*.entity{.ts,.js}'],
  synchronize: true,
  logging: true,
  migrations: ['dist/src/db/migrations/*.js'],
  cli: {
    migrationsDir: 'src/db/migrations',
  },
  timezone: 'Z',
};
export default ormConfig;
