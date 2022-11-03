import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { CartController } from './controllers/cart.controller';
import { Cart } from './entities/cart.entity';
import { CartRepository } from './repositories/cart.repository';
import { CartService } from './services/cart.service';

@Module({
  imports: [TypeOrmModule.forFeature([Cart])],
  controllers: [CartController],
  providers: [CartRepository, CartService],
  exports: [],
})
export class CartModule {}
