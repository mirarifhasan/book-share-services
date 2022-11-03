import { CreateDateColumn, UpdateDateColumn } from 'typeorm';

export class AppBaseEntity {
  @CreateDateColumn({ type: 'timestamp', select: false })
  created_at: Date;

  @UpdateDateColumn({ type: 'timestamp', select: false })
  updated_at: Date;
}
