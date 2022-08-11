import { HttpClient, HttpHeaders } from '@angular/common/http';
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
  getToken() {
    return localStorage.getItem("id_token")
  }

  private headers = new HttpHeaders().set('Content-Type', 'application/json')
  
  user!: User;
  isAuthenticated = false;

  constructor(private http: HttpClient) { }


  authWithToken(token: string) {
    let headers = new HttpHeaders().set("Authorization", token)
    console.log(headers);
    
    let observer = this.http.post<{message: string, id_token: string, expires_at: string}>(
      `${environment.serverUrl}/api/auth/check_token`, 
      {},
      {headers: headers, observe: 'response'}
    )

    observer.subscribe((response) => {
      if (response.status == 200 && response.body) {
        this.isAuthenticated = true
        this.setSession({ expiresIn: response.body.expires_at, idToken: response.body.id_token })
      }
    })

    return observer
  }

  register(username: string, email:string, password: string, password2: string) {
    let observer = this.http.post<{status: string, message: string, id_token: string, expires_at: string}>(
      `${environment.serverUrl}/api/auth/register`, 
      {username, email, password, password2}, 
      {headers: this.headers, observe: 'response'}
    )

    observer.subscribe((response) => {
      if (response.status == 200 && response.body) {
        this.isAuthenticated = true
        this.setSession({ expiresIn: response.body.expires_at, idToken: response.body.id_token })
      }
    })

    return observer
  }

  login(username: string, password: string) {

    let observer = this.http.post<{status: string, message: string, id_token: string, expires_at: string}>(
      `${environment.serverUrl}/api/auth/login`, 
      {username, password}, 
      {headers: this.headers, observe: 'response'}
    )

    observer.subscribe((response) => {
      if (response.status == 200 && response.body) {
        this.isAuthenticated = true
        this.setSession({ expiresIn: response.body.expires_at, idToken: response.body.id_token })
      } 
    })

    return observer
  }

  getProfile(id: number) {
    let observer = this.http.post<User>(
      `${environment.serverUrl}/api/users/profile`, 
      {id}, 
      {observe: 'response'}
    )

    observer.subscribe((response) => {
      if (response.status == 200)
        this.isAuthenticated = true   
      console.log(response.body?.username);
    })

    return observer
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


