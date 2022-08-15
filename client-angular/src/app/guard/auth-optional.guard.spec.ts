import { TestBed } from '@angular/core/testing';

import { AuthOptionalGuard } from './auth-optional.guard';

describe('AuthOptionalGuard', () => {
  let guard: AuthOptionalGuard;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    guard = TestBed.inject(AuthOptionalGuard);
  });

  it('should be created', () => {
    expect(guard).toBeTruthy();
  });
});
