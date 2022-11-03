import { Module } from '@nestjs/common';
import { CartModule } from './cart/cart.module';
import { ConfigModule } from '@nestjs/config';
import { DatabaseModule } from './common/database/database.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      envFilePath: '.env',
    }),
    DatabaseModule,
    CartModule,
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
