import { Body, Controller, Post } from '@nestjs/common';
import { CartService } from '../services/cart.service';

@Controller({ path: 'carts', version: '1' })
export class CartController {
  constructor(private cartService: CartService) {}

  @Post()
  async createCart(@Body() dto: { buyer_id: number; product_ids: number[] }) {
    let data = await this.cartService.createCart(dto);
    return { message: 'Cart created', data };
  }
}
