import { HttpClient, HttpHeaders } from '@angular/common/http';
import { ChangeDetectorRef, Injectable } from '@angular/core';
import { Observable, Subject, of, from } from 'rxjs';
import { tap, shareReplay, filter, map } from 'rxjs/operators';
import { User, UserCompressed } from 'src/models/user';
import * as moment from "moment";
import { environment } from 'src/environments/environment';
import { UserPage } from 'src/models/UserPage';

@Injectable()
export class AuthService {
  getToken() {
    return localStorage.getItem("id_token")
  }

  getTokenHeader() {
    const token = this.getToken()
    if (!token)
      return false
    
    return new HttpHeaders().set("Authorization", token)
  }

  private headers = new HttpHeaders().set('Content-Type', 'application/json')
  
  userSubscription?: Observable<User | false>;
  user?: User;
  isAuthenticated = true;

  constructor(private http: HttpClient) { }


  authWithToken(token: string) {
    let headers = new HttpHeaders().set("Authorization", token)
    
    let observer = this.http.post<User>(
      `${environment.serverUrl}/api/auth/check_token`, 
      {},
      {headers: headers, observe: 'response'}
    ).pipe(
      map((response) => {
        if (response.status == 200 && response.body && response.body.id_token) {
          this.isAuthenticated = true
  
          this.user = response.body
          
          this.setSession({ expiresIn: response.body.expires_at, idToken: response.body.id_token })
          return response.body
        } else {
          this.isAuthenticated = false
          this.logout()
          return false;
        }
      })
    )

    this.userSubscription = observer

    return observer
  }

  register(username: string, email:string, password: string, password2: string) {
    let observer = this.http.post<{status: string, message: string, id_token: string, expires_at: string, user_id: number}>(
      `${environment.serverUrl}/api/auth/register`, 
      {username, email, password, password2}, 
      {headers: this.headers, observe: 'response'}
    )

    return observer.pipe(
      map(response => {
        console.log(response);
        
        if (response.status == 200 && response.body) {
          this.isAuthenticated = true
          this.setSession({ expiresIn: response.body.expires_at, idToken: response.body.id_token })
        } 
        return response
      })
    )
  }

  login(username: string, password: string) {

    let observer = this.http.post<{status: string, message: string, id_token: string, expires_at: string, user_id: number}>(
      `${environment.serverUrl}/api/auth/login`, 
      {username, password}, 
      {headers: this.headers, observe: 'response'}
    )

    return observer.pipe(
      map(response => {
        if (response.status == 200 && response.body) {
          this.isAuthenticated = true
          this.setSession({ expiresIn: response.body.expires_at, idToken: response.body.id_token })
        } 
        return response
      })
    )
  }

  getCompressedProfile(id: number) {
    let subscription = this.getProfile(id)
    if (subscription == false)
      return false

    return subscription.pipe(
      map((response) => {
        console.log(response);
        if (response != false) {
          return response as UserCompressed
        }
        return false;
      })
    )
  }

  getProfile(id: number) {
    let headers = new HttpHeaders()

    const token = this.getToken()
    if (token) {
      headers = headers.set("Authorization", token)
    } else if (id == 0) {
      return false
    }

    let observer = this.http.post<UserPage>(
      `${environment.serverUrl}/api/users/profile`, 
      {id}, 
      {headers: headers, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        console.log(response);
        if (response.status == 200 && response.body) {
          return response.body as UserPage
        }
        return false;
      })
    )
  }

  changeProfile(name: string, birth_date: Date, city: string) {
    const token = this.getToken()
    console.log(token);
    
    if (!token)
      return false
    
    let headers = new HttpHeaders().set("Authorization", token)

    let observer = this.http.post<User>(
      `${environment.serverUrl}/api/users/change_profile`, 
      {name, birth_date, city}, 
      {headers: headers, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        console.log(response);
        
        if (response.status == 200 && response.body)
          return response.body as User
        return false;
      })
    )
  }

  changeProfileStatus(status: string) {
    const token = this.getToken()
    if (!token)
      return false
    
    let headers = new HttpHeaders().set("Authorization", token)
    

    let observer = this.http.post<User>(
      `${environment.serverUrl}/api/users/change_profile`, 
      {status}, 
      {headers: headers, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body)
          return response.body as User
        return false;
      })
    )
  }

  changeProfileAvatar(imageSrc: File) {
    const token = this.getToken()
    if (!token)
      return false
    
    const headers = new HttpHeaders({ "Authorization" : token, 'enctype': 'multipart/form-data' });
    
    var formData = new FormData();
    formData.append('avatar', imageSrc);
    console.log(imageSrc);
    

    let observer = this.http.post(
      `${environment.serverUrl}/api/users/change_avatar`, 
      formData, 
      {headers: headers, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body)
          return response.body
        return false;
      })
    )
  }

  getFriends(user_id: number) {
    let observer = this.http.get(
      `${environment.serverUrl}/api/users/get_friends/${user_id}`, 
      {observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          return response.body as UserCompressed[]
        }
        return false;
      })
    )
  }

  addToFriends(user_id: number) {
    const token = this.getToken()
    if (!token)
      return false
    
    let headers = new HttpHeaders().set("Authorization", token)

    let observer = this.http.post(
      `${environment.serverUrl}/api/users/add_to_friend`, 
      {user_id}, 
      {headers: headers, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          return response.body
        }
        return false;
      })
    )
  }

  deleteFriend(user_id: number) {
    const token = this.getToken()
    if (!token)
      return false
    
    let headers = new HttpHeaders().set("Authorization", token)

    let observer = this.http.post(
      `${environment.serverUrl}/api/users/delete_friend`, 
      {user_id}, 
      {headers: headers, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          return response.body
        }
        return false;
      })
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


