import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthService } from '../services/backend-api/auth.service';

@Injectable({
  providedIn: 'root'
})
export class AuthOptionalGuard implements CanActivate {
  constructor(private auth: AuthService) {}

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot){
      let token = this.auth.getToken();
      if (!token)
        return true;
      
      this.auth.authWithToken(token);
      return true
  }
  
}
