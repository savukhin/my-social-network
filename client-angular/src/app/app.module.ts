import { ChangeDetectorRef, Injectable, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HttpHandler, HttpInterceptor, HttpRequest, HTTP_INTERCEPTORS } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { UserPageComponent } from './profile/user-page/user-page.component';
import { AppRoutingModule } from './app-routing.module';
import { MessagesComponent } from './messages/messages.component';
import { ChatComponent } from './chat/chat.component';
import { FriendsComponent } from './friends/friends.component';
import { UpdateProfileComponent } from './profile/update-profile/update-profile.component';
import { LoginComponent } from './auth/login/login.component';
import { MainLayoutComponent } from './main-layout/main-layout.component';
import { RegisterComponent } from './auth/register/register.component';
import { ChangePasswordComponent } from './profile/change-password/change-password.component';
import { AuthService } from './services/backend-api/auth.service';

@Injectable()
export class XhrInterceptor implements HttpInterceptor {
  constructor(private auth: AuthService) {

  }

  intercept(req: HttpRequest<any>, next: HttpHandler) {
    let xhr = req.clone({
      headers: req.headers.set('X-Requested-With', 'XMLHttpRequest')
    });

    console.log(this.auth.isAuthenticated());

    // if (this.auth.isAuthenticated()) {
      console.log(this.auth.getToken());
      
      xhr = xhr.clone({
        headers: xhr.headers.set('X-Auth-Token', this.auth.getToken())
      });
    // }

    return next.handle(xhr);
  }
}

@NgModule({
  declarations: [
    AppComponent,
    UserPageComponent,
    MessagesComponent,
    ChatComponent,
    FriendsComponent,
    UpdateProfileComponent,
    LoginComponent,
    MainLayoutComponent,
    RegisterComponent,
    ChangePasswordComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    CommonModule,
    FormsModule,
    ReactiveFormsModule
  ],
  providers: [AuthService, { provide: HTTP_INTERCEPTORS, useClass: XhrInterceptor, multi: true }],
  bootstrap: [AppComponent]
})
export class AppModule { }
