import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
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
import { ContentService } from './services/backend-api/content.service';
import { WebsocketService } from './services/backend-api/websocket/websocket.service';
import { ChatService } from './services/backend-api/chat.service';

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
  providers: [AuthService, ContentService, ChatService],
  bootstrap: [AppComponent]
})
export class AppModule { }
