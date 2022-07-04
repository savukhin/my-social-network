import { Component, OnInit, ViewChild, ElementRef, AfterViewInit } from '@angular/core';

@Component({
    selector: 'app-user-page',
    templateUrl: './user-page.component.html',
    styleUrls: ['./user-page.component.scss']
})
export class UserPageComponent implements AfterViewInit {
    profile = {
        name: "Saveliy Karpukhin",
        isOnline: true,
        status: "#notforwar",
        birthDate: "30.08.2002",
        city: "Moscow",
    }
    editStatus = false
    
    // @ViewChild('status', { static: false, read: ElementRef }) status: ElementRef<HTMLInputElement> = {} as ElementRef;
    @ViewChild('pleaseDoIt') input: ElementRef<HTMLInputElement> = {} as ElementRef;;
    // @ViewChild('pleaseDoIt') input: ElementRef;

    showEditStatus() {
        // document.getElementById('statusInput')?.focus();
        setTimeout(() => {
            this.input.nativeElement.focus();
            this.input.nativeElement.value = this.profile.status;
        }, 0)
        this.editStatus = true;
    }

    hideEditStatus() {
        this.editStatus = false;
    }

    // constructor() { }
    ngAfterViewInit(): void {
        console.log(this.input.nativeElement.textContent)
    }
}
