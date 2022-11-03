import { Module } from '@nestjs/common';
import { CartController } from './controllers/cart.controller';

@Module({
  imports: [],
  controllers: [CartController],
  providers: [],
  exports: [],
})
export class CartModule {}
