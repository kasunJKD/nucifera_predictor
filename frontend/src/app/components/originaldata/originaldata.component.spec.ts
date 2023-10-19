import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OriginaldataComponent } from './originaldata.component';

describe('OriginaldataComponent', () => {
  let component: OriginaldataComponent;
  let fixture: ComponentFixture<OriginaldataComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ OriginaldataComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OriginaldataComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
