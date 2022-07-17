import { HttpClient } from '@angular/common/http';
import { ChangeDetectorRef, Injectable } from '@angular/core';
import { Observable, Subject, of, from } from 'rxjs';
import { tap, shareReplay, filter } from 'rxjs/operators';
import { User } from 'src/models/user';
import * as moment from "moment";
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private user = new Subject<User>();
  // user$ = this.user.asObservable();

  getUser(): Observable<User|undefined> {
    return this.user.asObservable().pipe(
      filter(user => !!user)
    );
  }

  constructor(private http: HttpClient, private cdref: ChangeDetectorRef) { }

  setUser(user: User) {
    this.user.next(user);
  }

  authWithToken(token: string) {
    let new_user = new User(0, "Saveliy Karpukhin");
    return of(new_user);
    // return new Observable<User>((observer) => {
    //   setTimeout(() => {
    //     observer.next(tmp)
    //   }, 1000)
    // })
  }

  register(login: string, email:string, password: string, password2: string) {
    return this.http.post<String>(`${environment.serverUrl}/user/registration`, {login, email, password, password2}).pipe(
      // tap((res: any) => this.setSession),
      shareReplay()
    )
    // return of("TOKEN");
  }

  login(login: string, password: string) {
    // return of("TOKEN");
    return this.http.post<User>(`${environment.serverUrl}/user/login`, {login, password}).pipe(
        // tap((res: any) => this.setSession),
        shareReplay()
      )
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
}


