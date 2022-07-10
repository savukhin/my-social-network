import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, Subject, of, from } from 'rxjs';
import { tap, shareReplay } from 'rxjs/operators';
import { User } from 'src/models/user';
import * as moment from "moment";

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private userLogged = new Subject<User>();
  userLogged$ = this.userLogged.asObservable();

  constructor(private http: HttpClient) { }

  register(login: string, email:string, password: string, password2: string) {
    return of("TOKEN");
  }

  login(login: string, password: string) {
    return of("TOKEN");
    // return this.http.post<User>('/api/auth/login', {login, password}).pipe(
    //     tap((res: any) => this.setSession),
    //     shareReplay()
    //   )
  }

  private setSession(authResult: { expiresIn: any; idToken: string; }) {
    const expiresAt = moment().add(authResult.expiresIn,'second');

    localStorage.setItem('id_token', authResult.idToken);
    localStorage.setItem("expires_at", JSON.stringify(expiresAt.valueOf()) );
  }

  logout() {
    localStorage.removeItem("id_token");
    localStorage.removeItem("expires_at");
  }

  public isLoggedIn() {
      return moment().isBefore(this.getExpiration());
  }

  isLoggedOut() {
      return !this.isLoggedIn();
  }

  getExpiration() {
      const expiration = localStorage.getItem("expires_at");
      const expiresAt = JSON.parse(expiration!);
      return moment(expiresAt);
  }    

  authWithToken(token: string) {
    // return of(new User(0, "Saveliy Karpukhin", true));
    return of(undefined);
  }
}


