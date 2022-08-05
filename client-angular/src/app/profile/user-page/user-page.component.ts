import { Component, OnInit, ViewChild, ElementRef, AfterViewInit, ChangeDetectorRef } from '@angular/core';
import { AuthService } from 'src/app/services/backend-api/auth.service';
import { User } from 'src/models/user';

@Component({
    selector: 'app-user-page',
    templateUrl: './user-page.component.html',
    styleUrls: ['./user-page.component.scss']
})
export class UserPageComponent implements AfterViewInit {
    user?: User;
    profile = new User(0, "savukhin", "Saveliy Karpukhin", true, "#notforwar", "30.08.2002", "Moscow")
    editStatus = false
    
    @ViewChild('pleaseDoIt') input: ElementRef<HTMLInputElement> = {} as ElementRef;

    showEditStatus() {
        setTimeout(() => {
            this.input.nativeElement.focus();
            this.input.nativeElement.value = this.profile.status;
        }, 0)
        this.editStatus = true;
    }

    hideEditStatus() {
        this.editStatus = false;
    }

    constructor(private auth: AuthService, private cdref: ChangeDetectorRef) {
        console.log("subscription");
        this.auth.getUser().subscribe(
            user => {
                console.log("UserPageComponent: User changed in userpage");

                if (user) {
                    this.profile = user;
                }
                this.cdref.detectChanges();
            }
        )
    }
    
    ngAfterViewInit(): void {
        this.auth.getProfile("savukhin").subscribe(
            data => {
                this.user = data;
            }
        )
    }
}
