#include<GL/gl.h>
#include<GL/glut.h>
#include <stdio.h>

int x1, y1, x2, y2;



void draw_pixel(int x, int y) {
	glBegin(GL_POINTS);
	glPointSize(3.0);
	glVertex2i(x, y);
	glEnd();
}

void draw_line(int x1, int y1, int x2, int y2) {
	int dx, dy, i, e;
	int incx, incy, inc1, inc2;
	int x,y;

	dx = x2-x1;
	dy = y2-y1;

	if (dx < 0) dx = -dx;
	if (dy < 0) dy = -dy;
	incx = 1;
	if (x2 < x1) incx = -1;
	incy = 1;
	if (y2 < y1) incy = -1;
	x = x1; y = y1;
	if (dx > dy) {
		draw_pixel(x, y);
		e = 2 * dy-dx;
		inc1 = 2*(dy-dx);
		inc2 = 2*dy;
		for (i=0; i<dx; i++) {
			if (e >= 0) {
				y += incy;
				e += inc1;
			}
			else
				e += inc2;
			x += incx;
			draw_pixel(x, y);
		}

	} else {
		draw_pixel(x, y);
		e = 2*dx-dy;
		inc1 = 2*(dx-dy);
		inc2 = 2*dx;
		for (i=0; i<dy; i++) {
			if (e >= 0) {
				x += incx;
				e += inc1;
			}
			else
				e += inc2;
			y += incy;
			draw_pixel(x, y);
		}
	}
glFlush();
}

void myDisplay(void) {
	//draw_line(x1, x2, y1, y2);
 glClear(GL_COLOR_BUFFER_BIT);
// bha//
  draw_line(125,150,125,300);
  draw_line(60,225,60,300);
  draw_line(60,225,125,225);
 
  //aa//
  draw_line(150,300,150,150);
//ga//
 draw_line(190,300,190,225);
  draw_line(250,300,250,150);
//
//wa//
  draw_line(320,300,320,150);
//
//upar ka line//
  draw_line(90,300,360,300);
// surname//
 draw_line(420,300,590,300);
//da//
  draw_line(460,300,460,260);
  draw_line(485,300,485,150);
    draw_line(460,260,460,260);
   //draw_line(485,150,570,280);

//sa
   draw_line(570,300,570,150);
   
   //draw_line(570,225,515,225);
  draw_line(570,225,515,225);
   draw_line(515,225,540,150);
	
double xu = 0.0 , yu = 0.0 , u = 0.0 ; 
    int i = 0 ; 
    for(u = 0.0 ; u <= 1.0 ; u += 0.0001) 
    { 
        xu = (1-u)*(1-u)*(1-u)*253+3*(1-u)*(1-u)*u*235+3*(1-u)*u*u*235 
             +250*u*u*u; 
       yu = (1-u)*(1-u)*(1-u)*355+3*(1-u)*(1-u)*u*336+3*(1-u)*u*u*317 
             +300*u*u*u; 
draw_pixel(xu, yu);
  }

double xu1 = 0.0 , yu1 = 0.0 , u1 = 0.0 ; 
    int i1 = 0 ; 
    for(u1 = 0.0 ; u1 <= 1.0 ; u1 += 0.0001) 
    { 
        xu1 = (1-u1)*(1-u1)*(1-u1)*60+3*(1-u1)*(1-u1)*u1*45+3*(1-u1)*u1*u1*45 
             +60*u1*u1*u1; 
       yu1 = (1-u1)*(1-u1)*(1-u1)*300+3*(1-u1)*(1-u1)*u1*300+3*(1-u1)*u1*u1*287
             +287*u1*u1*u1; 
draw_pixel(xu1, yu1);


  }
draw_line(60,225,60,212);

double xu2 = 0.0 , yu2 = 0.0 , u2 = 0.0 ; 
    int i2 = 0 ; 
    for(u2 = 0.0 ; u2 <= 1.0 ; u2 += 0.0001) 
    { 
        xu2 = (1-u2)*(1-u2)*(1-u2)*60+3*(1-u2)*(1-u2)*u2*45+3*(1-u2)*u2*u2*45 
             +60*u2*u2*u2; 
       yu2 = (1-u2)*(1-u2)*(1-u2)*225+3*(1-u2)*(1-u2)*u2*225+3*(1-u2)*u2*u2*225
             +212*u2*u2*u2; 
draw_pixel(xu2, yu2);


	}

draw_line(190,300,190,212);

double xu3 = 0.0 , yu3 = 0.0 , u3 = 0.0 ; 
    int i3 = 0 ; 
    for(u3 = 0.0 ; u3 <= 1.0 ; u3 += 0.0001) 
    { 
        xu3 = (1-u3)*(1-u3)*(1-u3)*190+3*(1-u3)*(1-u3)*u3*175+3*(1-u3)*u3*u3*175 
             +190*u3*u3*u3; 
       yu3 = (1-u3)*(1-u3)*(1-u3)*225+3*(1-u3)*(1-u3)*u3*225+3*(1-u3)*u3*u3*225
             +212*u3*u3*u3; 
draw_pixel(xu3, yu3);


	}

double xu4 = 0.0 , yu4 = 0.0 , u4 = 0.0 ; 
    int i4 = 0 ; 
    for(u4 = 0.0 ; u4 <= 1.0 ; u4 += 0.0001) 
    { 
        xu4 = (1-u4)*(1-u4)*(1-u4)*320+3*(1-u4)*(1-u4)*u4*265+3*(1-u4)*u4*u4*265
             +320*u4*u4*u4; 
       yu4 = (1-u4)*(1-u4)*(1-u4)*250+3*(1-u4)*(1-u4)*u4*275+3*(1-u4)*u4*u4*175
             +200*u4*u4*u4; 
draw_pixel(xu4, yu4);
	}

	double xu5 = 0.0 , yu5 = 0.0 , u5 = 0.0 ; 
    int i5 = 0 ; 
    for(u5 = 0.0 ; u5 <= 1.0 ; u5 += 0.0001) 
    { 
        xu5 = (1-u5)*(1-u5)*(1-u5)*515+3*(1-u5)*(1-u5)*u5*550+3*(1-u5)*u5*u5*527
             +515*u5*u5*u5; 
       yu5 = (1-u5)*(1-u5)*(1-u5)*299+3*(1-u5)*(1-u5)*u5*301+3*(1-u5)*u5*u5*225
             +225*u5*u5*u5; 
draw_pixel(xu5, yu5);

	}
	double xu6 = 0.0 , yu6 = 0.0 , u6 = 0.0 ; 
    int i6 = 0 ; 
    for(u6 = 0.0 ; u6 <= 1.0 ; u6 += 0.0001) 
    { 
        xu6 = (1-u6)*(1-u6)*(1-u6)*460+3*(1-u6)*(1-u6)*u6*390+3*(1-u6)*u6*u6*390
             +470*u6*u6*u6; 
       yu6 = (1-u6)*(1-u6)*(1-u6)*260+3*(1-u6)*(1-u6)*u6*260+3*(1-u6)*u6*u6*150
             +150*u6*u6*u6; 
draw_pixel(xu6, yu6);
	}
	double xu7 = 0.0 , yu7 = 0.0 , u7 = 0.0 ; 
    int i7 = 0 ; 
    for(u7 = 0.0 ; u7 <= 1.0 ; u7 += 0.0001) 
    { 
        xu7 = (1-u7)*(1-u7)*(1-u7)*450+3*(1-u7)*(1-u7)*u7*450+3*(1-u7)*u7*u7*470
             +470*u7*u7*u7; 
       yu7 = (1-u7)*(1-u7)*(1-u7)*150+3*(1-u7)*(1-u7)*u7*170+3*(1-u7)*u7*u7*170
             +150*u7*u7*u7; 
draw_pixel(xu7, yu7);
	}
 draw_line(450,150,465,124);

glFlush();
}

void myInit() {
	glClear(GL_COLOR_BUFFER_BIT);
		glPointSize(4.0);
	glClearColor(0.0, 0.0, 0.0, 0.0);
	 glColor3f(3.0,1.0,0.0);
        gluOrtho2D(0,640,0,480);
  }

int main(int argc, char **argv) {
	glutInit(&argc, argv);
	glutInitDisplayMode(GLUT_SINGLE|GLUT_RGB);
	glutInitWindowSize(1366, 768);
	glutInitWindowPosition(0, 0);
	glutCreateWindow("Name in Devanagari-Script");
		glPointSize(4.0);
	myInit();
	glutDisplayFunc(myDisplay);
	glutMainLoop();
        return 0;
}
