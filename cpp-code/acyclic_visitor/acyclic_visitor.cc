#include <iostream>

// Forward declaration of shape types for the visitor
class Circle;
class Square;

// Base Visitor class for shapes
class ShapeVisitor {
public:
    virtual void visitCircle(Circle* circle) {}
    virtual void visitSquare(Square* square) {}
    virtual ~ShapeVisitor() {}
};

// Base Shape class
class Shape {
public:
    virtual void accept(ShapeVisitor& visitor) = 0;
    virtual ~Shape() {}
};

// Circle shape class
class Circle : public Shape {
public:
    void accept(ShapeVisitor& visitor) override {
        visitor.visitCircle(this);
    }
    // Circle-specific methods
};

// Square shape class
class Square : public Shape {
public:
    void accept(ShapeVisitor& visitor) override {
        visitor.visitSquare(this);
    }
    // Square-specific methods
};

// Concrete Visitor class for rendering shapes
class RenderVisitor : public ShapeVisitor {
public:
    void visitCircle(Circle* circle) override {
        std::cout << "Rendering a Circle." << std::endl;
    }
    void visitSquare(Square* square) override {
        std::cout << "Rendering a Square." << std::endl;
    }
};

int main() {
    Circle circle;
    Square square;
    RenderVisitor renderVisitor;

    circle.accept(renderVisitor); // Output: Rendering a Circle.
    square.accept(renderVisitor); // Output: Rendering a Square.

    return 0;
}


