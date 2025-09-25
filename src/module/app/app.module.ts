import { Module } from '@nestjs/common';
import { AppController } from '../../interface/controller/api/rest/app.controller';
import { AppService } from '../../service/app.service';

@Module({
  imports: [],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
