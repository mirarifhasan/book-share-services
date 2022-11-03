import { Injectable } from '@nestjs/common';
import { CartRepository } from '../repositories/cart.repository';

@Injectable()
export class CartService {
  constructor(private cartRepository: CartRepository) {}

  async createCart(dto: { buyer_id: number; product_ids: number[] }) {
    for (let i = 0; i < dto.product_ids.length; i++) {
      await this.cartRepository.save({
        buyer_id: dto.buyer_id,
        product_id: dto.product_ids[i],
      });
    }
    return null;
  }
}
