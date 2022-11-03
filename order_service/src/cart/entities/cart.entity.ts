import { AppBaseEntity } from 'src/common/database/entity/app-base.entity';
import { Entity, PrimaryGeneratedColumn, Column } from 'typeorm';
@Entity({
  name: 'carts',
})
export class Cart extends AppBaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  buyer_id: number;

  @Column()
  product_id: number;
}
