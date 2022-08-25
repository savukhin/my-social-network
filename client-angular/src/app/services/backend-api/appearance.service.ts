import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable()
export class AppearanceService {
  public blackout = new Subject<boolean>()

  constructor() { }
}
