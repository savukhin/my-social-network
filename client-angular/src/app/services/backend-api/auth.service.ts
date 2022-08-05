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

  private user = new Subject<User>();
  // user$ = this.user.asObservable();
  authenticated = false;
  token = "";

  isAuthenticated(): boolean {
    // return this.authenticated;
    return localStorage.getItem("authenticated")! == "true";
  }

  getToken(): string {
    // return this.token;
    return localStorage.getItem("token")!;
  }

  getUser(): Observable<User|undefined> {
    return this.user.asObservable().pipe(
      filter(user => !!user)
    );
  }

  constructor(private http: HttpClient) { }

  setUser(user: User) {
    this.user.next(user);
  }
  
  updateToken() {
    return this.http.get(`${environment.serverUrl}/user/token`, { headers: {}}).subscribe(( data ) => {
      this.token = (data as { token: string }).token;
      console.log(data);
      localStorage.setItem('token', this.token);
    });
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

  register(login: string, email:string, password: string, password2: string, callback = () => {}) {
    return this.http.post<{message: string}>(`${environment.serverUrl}/user/registration`, {username: login, email, password, password2})
      .subscribe(response => {
        console.log(response['message']);
        
        // if (response['message'] == "OK") {
        this.authenticated = true;
        localStorage.setItem('authenticated', "true");
        // } else {
        //   this.authenticated = false;
        // }
        console.log(response);
        console.log(this.authenticated);
        
        

        // authWithToken();
        this.updateToken();
        // return this.login(login, password, callback)
        return callback && callback();
      })
  }

  login(userName: string, password: string, callback = () => {}) {
    const credentials = {
      userName: userName,
      password: password
    }

    const headers = new HttpHeaders(credentials ? {
      authorization : 'Basic ' + btoa(credentials.userName + ':' + credentials.password)
    } : {auth: "empty"});

    console.log(headers);
    console.log(credentials);
    
    return this.http.get<User>(`${environment.serverUrl}/user/login`, { headers: headers })
      .subscribe(response => {
        if (response['name']) {
          this.authenticated = true;
        } else {
          this.authenticated = false;
        }
        
        this.updateToken();
        return callback && callback();
      })
  }

  getProfile(username: string, callback = () => {}) {
    return this.http.get<User>(`${environment.serverUrl}/user/profile/${username}`);
      // .subscribe(response => {
      //   callback && callback();
      //   return response;
      // })
  }

  private setSession(authResult: { expiresIn: any; idToken: string; }) {
    const expiresAt = moment().add(authResult.expiresIn,'second');

    localStorage.setItem('id_token', authResult.idToken);
    localStorage.setItem("expires_at", JSON.stringify(expiresAt.valueOf()) );
  }

  logout() {
    // localStorage.removeItem("id_token");
    // localStorage.removeItem("expires_at");
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


