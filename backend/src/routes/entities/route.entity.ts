import { Prop, raw, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type RouteDocument = Route & Document;

type Position = { lat: number; long: number };

@Schema()
export class Route {
  @Prop()
  title: string;

  @Prop(
    raw({
      lat: { type: Number },
      long: { type: Number },
    }),
  )
  startPosition: Position;

  @Prop(
    raw({
      lat: { type: Number },
      long: { type: Number },
    }),
  )
  endPosition: Position;
}

export const RouteSchema = SchemaFactory.createForClass(Route);
